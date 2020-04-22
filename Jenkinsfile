pipeline {
  agent {
    kubernetes {
        label 'go-pipeline-pod'
        yamlFile 'podTemplate/go-pipeline-pod.yaml'
        idleMinutes 120
    }
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
            sh 'gcloud auth activate-service-account goci-example-jefff@soldev-dev.iam.gserviceaccount.com --key-file=${GC_KEY}'
            sh 'gcloud container clusters get-credentials staging --zone us-central1-c --project soldev-dev'
            sh 'helm install goci-example chart/goci-example'
          }
        }
      }
    }
    stage('Staging Test') {
       steps {
           container('golang'){
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
}