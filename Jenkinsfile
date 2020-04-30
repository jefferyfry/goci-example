pipeline {
  agent {
    kubernetes {
        label 'go-pipeline-pod'
        yamlFile 'podTemplate/go-pipeline-pod.yaml'
        idleMinutes 120
    }
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
            sh "docker build -t partnership-public-images.jfrog.io/goci-example:$BUILD_NUMBER ."
        }
      }
    }
    stage('Docker Push to Staging Repo') {
      steps {
        container('jfrog-cli-go'){
            withCredentials([usernamePassword(credentialsId: 'stagingrepo', passwordVariable: 'APIKEY', usernameVariable: 'USER')]) {
                sh "jfrog rt dp partnership-public-images.jfrog.io/goci-example:$BUILD_NUMBER public-images --user=$USER --apiKey=$APIKEY --build-name=$JOB_NAME --build-number=$BUILD_NUMBER"
            }
        }
      }
    }
  }
  post {
      success {
        script {
           sh "curl -v -XPOST -H \"authorization: Basic amVmZmY6amZyMGdqM25rMW5z\" \"https://partnership-pipelines-api.jfrog.io/v1/projectIntegrations/17/hook\" -d '{\"buildName\":\"$JOB_NAME\",\"buildNumber\":\"$BUILD_NUMBER\",\"buildInfoResourceName\":\"jenkinsBuildInfo\"}' -H \"Content-Type: application/json\""
        }
      }
  }
}