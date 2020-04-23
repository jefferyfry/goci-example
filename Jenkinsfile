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
          sh "docker build -t partnership-public-images.jfrog.io/staging/goci-example:${env.BUILD_NUMBER} ."
        }
      }
    }
    stage('Docker Push to Staging Repo') {
      steps {
        container('docker'){
            script {
              docker.withRegistry( 'https://partnership-public-images.jfrog.io', 'stagingrepo' ) {
                sh "docker push partnership-public-images.jfrog.io/staging/goci-example:${env.BUILD_NUMBER}"
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
            sh "helm install --set image.tag=${env.BUILD_NUMBER} --name goci-example --namespace staging ./chart/goci-example/"
          }
        }
      }
    }
    stage('Wait for Server') {
       steps {
          timeout(time: 1, unit: 'MINUTES') {
              waitUntil {
                script {
                    try {
                        def r = sh script: "curl -s --head  --request GET  ${env.STAGING_URL}/status | grep '200'", returnStdout: true
                        return (r == 'HTTP/1.1 200 OK');
                    } catch(Exception e){
                        return false;
                    }
                }
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
                     sh "docker tag partnership-public-images.jfrog.io/staging/goci-example::${env.BUILD_NUMBER} partnership-public-images.jfrog.io/release/goci-example::${env.BUILD_NUMBER}"
                     sh "docker push partnership-public-images.jfrog.io/release/goci-example::${env.BUILD_NUMBER}"
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