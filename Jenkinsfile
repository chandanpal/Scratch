pipeline {
  agent {
    node {
      label 'Build docker Image'
    }
  }
  stages {
    stage('Build vcenter') {
      steps {
        sh 'docker build -t dockersamples/vcenter ./ecoScripts/vcenter/Image'
      }
    } 
   
    stage('Push vcenter image') {
      when {
        branch 'Runner2.0'
      }
      steps {
        withDockerRegistry(credentialsId: 'dockerhub', url:'https://hub.docker.com') {
          sh 'docker push dockersamples/vcenter'
        }
      }
    }
  }
}
