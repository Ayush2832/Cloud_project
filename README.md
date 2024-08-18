# Deployed a Scalable Golang Web Application on Azure
### This is a simple cloud project where we create a docker image of our golang web-server and then create a docker image from it. After that we simply run this image in VMSS in azure and then finally add the load balancer with it.
### Steps of the project.
![image](https://github.com/user-attachments/assets/e433cf70-9ddf-42bd-b81e-76f8d46571c3)

1. Firt we created a simple golang web server. For that we use net/http package.
```go
package main
import (
    "io"
    "net/http"
)
func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/hello", getHello)
    err := http.ListenAndServe(":3333", nil)
    checkerr(err)
}
func checkerr(err error) {
    panic("unimplemented")
}
func getRoot(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "This is a website\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello HTTP\n")
}
```

2. Then we will create a simple docker file from it.
```dockerfile
FROM golang:1.21
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLE=0 GOOS=linux go build -o /server
EXPOSE 3333
CMD ["/server"]
```

3. After that we will run the container locally. Then we push the image to docker registry which is docker hub 
> `docker pull ayush2832/web-server` - image name

4. After that we will make a script to install the docker and run the continaer in azure VM. To run this scipt we go the advance option -> user data and paste the scipt.
```sh
#!/bin/bash
# Update package index and install required dependencies

sudo apt update
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

# Add Docker's official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add Docker repository
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io

# Start and enable Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Pull image and run container
docker pull ayush2832/web-server
IMAGE_ID=$(docker images -q ayush2832/web-server)
docker run -d --name container -p 5555:3333 "$IMAGE_ID"
```

![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240817142606.png)

5. After creating the container we will go to the Network security group and we will make a new inbound rule for the port 5555, becuase our container is working on port 5555.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818114909.png)

6. Then our container will start to run. Just type the IP address on the browser and add the port 5555 with it.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818115301.png)

7. After that we will create a Virtual Machine scale set, where we will create two virtual machine. After that we will similarly give the inbound rule for port 5555.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818120829.png)
Then our vm in Virtual Machine Scale set start to work.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818120808.png)

8. Finally we will create a load balancer. In load balancer we need to do some configuration
	- First we will choose the standard SKU
	- After that we will do FrontEnd ip configuration which will give the IP address to our load balancer.
	![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818125503.png)
	
	- After that we will create a backend pool. If your SKU of load balancer is same as SKU of the VMs then you can simply connect the VMs with the backend pool.
	![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818125215.png)
	
	- After that we need to set the load balancing rule. Here we need to define that the our frontend ip port will redirect the traffic to the backend poll vm port. Here my forntend ip port is 8080 and backend port is 5555.
        ![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818125503.png)
	
	- At last we will create a health probe and then simply make the load balancer
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818125531.png)

9. Simply create the load balancer.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818124930.png)

10. At last when you type `http://load_balancer_ip:8080`. Then it will simply show the result.
![img](https://github.com/Ayush2832/Cloud_project/blob/master/Pasted%20image%2020240818155940.png)
