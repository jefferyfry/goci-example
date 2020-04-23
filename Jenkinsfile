pipeline {
  agent {
    kubernetes {
        label 'go-pipeline-pod'
        yamlFile 'podTemplate/go-pipeline-pod.yaml'
        idleMinutes 120
    }
  }
  environment {
     STAGING_URL = 'http://goci-example.35.193.183.84.xip.io'
  }
  stages {
    stage('Compile') {
        steps {
            container('golang'){
                sh 'go build'
            }
        }
    }
    stage('Unit Tests') {
        steps {
            container('golang'){
                sh 'go test ./... -run Unit'
            }
        }
    }
    stage('Docker Build') {
      steps {
        container('docker'){
          sh 'docker build -t partnership-public-images.jfrog.io/staging/goci-example:latest .'
        }
      }
    }
    stage('Docker Push to Staging Repo') {
      steps {
        container('docker'){
            script {
              docker.withRegistry( 'https://partnership-public-images.jfrog.io', 'stagingrepo' ) {
                sh 'docker push partnership-public-images.jfrog.io/staging/goci-example:latest'
              }
           }
        }
      }
    }
    stage('Deploy to Staging') {
      steps {
        container('gcloud-kubectl-helm'){
          withCredentials([file(credentialsId: 'key-sa', variable: 'GC_KEY')]) {
            echo "Activating service account ${env.GC_KEY}"
            sh "gcloud auth activate-service-account --key-file=${env.GC_KEY}"
            sh 'gcloud config set project soldev-dev'
            sh 'gcloud container clusters get-credentials staging --zone us-central1-c --project soldev-dev'
            sh 'helm install --name goci-example --namespace staging ./chart/goci-example/'
          }
        }
      }
    }
    stage('Wait for Server') {
       steps {
          timeout(time: 60, unit: 'SECONDS') {
              waitUntil {
                sh "curl -s --head  --request GET  ${env.STAGING_URL} | grep '200'"
              }
          }
       }
    }
    stage('Staging Test') {
       steps {
           container('golang'){
              echo "Running staging tests against ${env.STAGING_URL}"
              sh 'go test ./... -run Staging'
           }
       }
    }
    stage('Docker Push to Release Repo') {
      steps {
        container('docker'){
            script {
               docker.withRegistry( 'https://partnership-public-images.jfrog.io', 'releaserepo' ) {
                     sh 'docker tag partnership-public-images.jfrog.io/staging/goci-example:latest partnership-public-images.jfrog.io/release/goci-example:latest'
                     sh 'docker push partnership-public-images.jfrog.io/release/goci-example:latest'
              }
            }
        }
      }
    }
  }
  post {
      success {
          sh 'helm delete --purge goci-example'
      }
  }
}