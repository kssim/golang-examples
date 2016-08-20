# How to set the environment for go-language.
I write here how to set the environment for build a go programs.<br>
But, the target is only linux based system not included windows system.
So, if you want to set the windows development environment, check this website [Install the Go tools](https://golang.org/doc/install#install).



# Setting the environment.
1. First you get the go language to [Golang.org](https://golang.org/dl/).
2. Second, you decompress the go language file to <code>/usr/local</code>.
	* <code>tar xvfz go1.7.linux-amd64.tar.gz -C /usr/local</code>.
3. Configured 'PATH' variable.
	* Add the go binary path to your profiles.
		* ex) <code>vi ~/.bashrc</code> and write <code>export PATH=$PATH:/usr/local/go/bin</code>.
	* Check the 'PATH' variable is worked.
		* Input the <code>go version</code> to shell, you can show the result <code>go version go1.7 linux/amd64</code>
