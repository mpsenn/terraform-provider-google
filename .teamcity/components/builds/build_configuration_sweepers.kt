package builds

import DefaultBuildTimeoutDuration
import DefaultParallelism
import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.sharedResources
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import replaceCharsId


fun BuildConfigurationForSweeper(providerName: String, sweeperName: String, packages: Map<String, Map<String, String>>, parentProjectName: String, vcsRoot: GitVcsRoot, sharedResources: List<String>, environmentVariables: AccTestConfiguration): BuildType {
    val sweeperPackage: Map<String, String> = packages.getValue("sweeper")
    val sweeperPath: String = sweeperPackage.getValue("path").toString()
    val s = SweeperDetails(sweeperName, parentProjectName, providerName)

    return s.sweeperBuildConfig(sweeperPath, vcsRoot, sharedResources, DefaultParallelism, environmentVariables)
}

class SweeperDetails(private val sweeperName: String, private val parentProjectName: String, private val providerName: String) {

    fun sweeperBuildConfig(
        path: String,
        vcsRoot: GitVcsRoot,
        sharedResources: List<String>,
        parallelism: Int,
        environmentVariables: AccTestConfiguration,
        buildTimeout: Int = DefaultBuildTimeoutDuration
    ): BuildType {

        // These hardcoded values affect the sweeper CLI command's behaviour
        val testPrefix = "TestAcc"
        val testTimeout = "12"
        val sweeperRegions = "us-central1"
        val sweeperRun = "" // Empty string means all sweepers run

        return BuildType {

            id(uniqueID())

            name = sweeperName

            vcs {
                root(vcsRoot)
                cleanCheckout = true
            }

            steps {
                // Commenting out these steps during refactoring the TeamCity config means we don't interact with the
                // GCP test projects while testing the new config
//                setGitCommitBuildId()
//                tagBuildToIndicatePurpose()
//                configureGoEnv()
//                downloadTerraformBinary()
//                runSweepers(sweeperName)
                helloWorld(sharedResources[0])
            }

            features {
                golang()
                if (sharedResources.isNotEmpty()) {
                    sharedResources {
                        // When the build runs, it locks the value(s) below
                        sharedResources.forEach { lock ->
                            lockAllPackageValues(lock)
                        }
                    }
                }
            }

            params {
                configureGoogleSpecificTestParameters(environmentVariables)
                terraformAcceptanceTestParameters(parallelism, testPrefix, testTimeout)
                terraformSweeperParameters(sweeperRegions, sweeperRun)
                terraformSkipProjectSweeper()
                terraformLoggingParameters(providerName)
                terraformCoreBinaryTesting()
                terraformShouldPanicForSchemaErrors()
                readOnlySettings()
                workingDirectory(path)
            }

            artifactRules = "%teamcity.build.checkoutDir%/debug*.txt"

            failureConditions {
                errorMessage = true
                executionTimeoutMin = buildTimeout
            }

        }
    }

    private fun uniqueID(): String {
        // Replacing chars can be necessary, due to limitations on IDs
        // "ID should start with a latin letter and contain only latin letters, digits and underscores (at most 225 characters)."
        var id = "%s_%s".format(this.parentProjectName, this.sweeperName)
        return replaceCharsId(id)
    }
}
