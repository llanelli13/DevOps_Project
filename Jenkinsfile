pipeline {
    agent any
    stages {
        stage('Clone Repository') {
            steps {
                git branch:"main", url:"https://github.com/llanelli13/DevOps_Project.git"
            }
        }
        stage('Build Application') {
            steps {
                script {
                    docker.image('golang:1.19').inside {
                        sh 'go mod tidy'
                        sh 'go build -o myapp .'
                    }
                }
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    def appImage = docker.build("myteam/myapp:latest")
                }
            }
        }
        stage('Run Container') {
            steps {
                script {
                    docker.run("myteam/myapp:latest", "-p 8080:8080")
                }
            }
        }
    }
}
