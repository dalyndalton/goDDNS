pipeline {
    agent any

    environment {
        CREDENTIALS = credentials('ddns-config')
        PATH = "${PATH}:/usr/local/go/bin"
    }

    triggers {
        cron('H/20 0 * * *') // Triggers every 20 minutes
    }

    stages {
        stage('Checkout & Build') {
            steps {
                checkout scm
                sh 'go build'
                sh 'chmod +x ddns-updater'
                sh 'ls'
            }
        }
        stage('Load Config') {
            steps {
                sh 'cp $CREDENTIALS config.yml'
            }
        }
        stage('Update DDNS') {
            steps {
                sh './ddns-updater config.yml'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
