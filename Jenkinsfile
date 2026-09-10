pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('API') {
            steps {
                dir('apps/api') {
                    sh 'go test ./...'
                    sh 'go build ./...'
                }
            }
        }

        stage('Web') {
            steps {
                dir('apps/web') {
                    sh 'pnpm install --frozen-lockfile'
                    sh 'pnpm lint'
                    sh 'pnpm build'
                }
            }
        }
    }
}
