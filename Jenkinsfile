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
    stage('Build') {
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
          sh "docker build -t partnership-public-images.jfrog.io/jenkins/staging/goci-example:${env.BUILD_NUMBER} ."
        }
      }
    }
    stage('Docker Push to Staging Repo') {
      steps {
        container('docker'){
            script {
              docker.withRegistry( 'https://partnership-public-images.jfrog.io', 'stagingrepo' ) {
                sh "docker push partnership-public-images.jfrog.io/jenkins/staging/goci-example:${env.BUILD_NUMBER}"
              }
           }
        }
      }
    }
  }
  post {
      success {
        script {
           sh "curl -XPOST -H \"authorization: Basic amVmZmY6amZyMGdqM25rMW5z\" \"https://partnership-pipelines-api.jfrog.io/v1/projectIntegrations/17/hook\" -d '{\"buildName\":\"gociexample_build\",\"buildInfoResourceName\":\"jenkinsBuildInfo\"}' -H \"Content-Type: application/json\""
        }
      }
  }
}