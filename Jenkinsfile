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
        container('golang'){
            steps {
                sh 'go build'
            }
        }
    }
    stage('Code Analysis') {
        steps {
            sh 'curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $GOPATH/bin v1.12.5'
            sh 'golangci-lint run'
        }
    }
    stage('Unit Tests') {
        container('golang'){
            steps {
                sh 'go test -run Unit'
            }
        }
    }
    stage('Docker Build') {
      steps {
        container('docker'){
          sh 'docker build -t partnership-jfrog-artifactory.jfrog.io/staging/goci-example:latest .'
        }
      }
    }
    stage('Docker Push to Staging Repo') {
      steps {
        container('docker'){
          withCredentials([usernamePassword(credentialsId: 'stagingrepo', usernameVariable: 'stagingrepouser', passwordVariable: 'stagingrepopassword')]) {
            sh 'docker login -u ${env.stagingrepouser} -p ${env.stagingrepopassword}'
            sh 'docker push partnership-jfrog-artifactory.jfrog.io/staging/goci-example:latest'
          }
        }
      }
    }
    stage('Deploy to Staging') {
      steps {
        container('gcloud-kubectl-helm'){
          withCredentials([file(credentialsId: 'key-sa', variable: 'GC_KEY')]) {
            sh 'gcloud auth activate-service-account --key-file=${GC_KEY}'
            sh 'gcloud container clusters get-credentials staging --zone us-central1-c --project soldev-dev'
            sh 'helm install goci-example chart/goci-example'
          }
        }
      }
    }
    stage('Staging Test') {
      container('golang'){
          steps {
              sh 'go test -run Staging'
          }
      }
    }
    stage('Docker Push to Release Repo') {
      steps {
        container('docker'){
          withCredentials([usernamePassword(credentialsId: 'releaserepo', usernameVariable: 'releaserepouser', passwordVariable: 'releaserepopassword')]) {
            sh "docker login -u ${env.releaserepouser} -p ${env.releaserepopassword}"
            sh 'docker push partnership-jfrog-artifactory.jfrog.io/release/goci-example:latest'
          }
        }
      }
    }
  }
}