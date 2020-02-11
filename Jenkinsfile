pipeline {
  agent any
  stages {
    stage('Build servicenow consumer Image') {
      steps {
        sh 'docker build -t sushmithakj2018/service_now-consumer ./ecoScripts/consumer/service_now-consumer/Image'
      }
    }
    stage('Build servicenow producer Image') {
      steps {
        sh 'docker build -t sushmithakj2018/service_now-producer ./ecoScripts/producer/service_now-producer/Image'
      }
    } 

   
    stage('Push servicenow consumer image') {
      steps {
        withDockerRegistry([credentialsId: 'dockerhub', url:'']) {
          sh 'docker push sushmithakj2018/service_now-consumer'
        }
      }
    }
    stage('Push servicenow producer image') {
      steps {
        withDockerRegistry([credentialsId: 'dockerhub', url:'']) {
          sh 'docker push sushmithakj2018/service_now-producer'
        }
      }
    }

  }
}
