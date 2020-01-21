pipeline {
  agent any
  stages {
    stage('Build vcenter') {
      steps {
        sh 'docker build -t dockersamples/vcenter ./ecoScripts/application/vcenter/Image'
      }
    } 
   
    stage('Push vcenter image') {
      when {
        branch 'master'
      }
      steps {
        withDockerRegistry(credentialsId: 'dockerhub', url:'https://hub.docker.com') {
          sh 'docker push dockersamples/vcenter'
        }
      }
    }
  }
}
