FROM centos:centos7.4.1708
LABEL Project=ecoHub Name=Target_Connectivity Version=0.0.1
WORKDIR /app
ADD . /app

# install pip and then install dependencies
RUN curl "https://bootstrap.pypa.io/get-pip.py" -o "get-pip.py"
RUN python get-pip.py
RUN pip install --trusted-host pypi.python.org -r requirements.txt

# run the app when the container launches
# ENTRYPOINT [ "/bin/bash", "eco_action.sh" ]
ENTRYPOINT [ "python", "eco_action.py" ]