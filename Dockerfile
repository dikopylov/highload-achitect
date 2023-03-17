# Specifies a parent image
FROM golang:1.19.2-bullseye
 
# Creates an app directory to hold your app’s source code
WORKDIR /opt/app
 
# Copies everything from your root directory into /opt/app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build -o /opt/app/bin/ /opt/app/cmd/social-network
 
# Specifies the executable command that runs when the container starts
CMD [ "/opt/app/bin/social-network" ]