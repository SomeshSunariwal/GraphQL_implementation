dockerImage = ''

pipeline{

    agent any

    environment {
        registry = "someshdokerbox/test"
        registryCredential = 'docker-hub'
    }

    stages{
        stgage("Build"){
            steps {
                script {
                    sh "make build"
                }
            }
        }
        stage("0Build Docker Image") {
            steps {
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }
        stage("Push Image") {
            steps {
                script {
                        docker.withRegistry( '', registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }
    }
}