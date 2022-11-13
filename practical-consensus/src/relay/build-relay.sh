cd /tmp
git clone https://github.com/amark/gun.git
cd /tmp/gun
sudo docker build -t in-browser-pow/gundb:v1 .
# sudo docker run -d -p 8765:8765 in-browser-pow/gundb:v1