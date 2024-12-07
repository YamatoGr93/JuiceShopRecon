# JuiceShopRecon

**JuiceShopRecon** is a Dockerized reconnaissance tool designed to test web applications for vulnerabilities. It performs DNS mapping, endpoint fuzzing, and port scanning, generating comprehensive reports to assist in identifying potential security issues. This project was specifically built to analyze the OWASP Juice Shop application and can be adapted for other targets.

---

## **Features**
1. **DNS Mapping**:
   - Extracts A, CNAME, MX, and TXT records for a given domain.
2. **Endpoint Fuzzing**:
   - Tests web application endpoints with predefined payloads to identify potential vulnerabilities.
3. **Port Scanning**:
   - Scans for open ports and attempts to identify services running on them.
4. **Comprehensive Reporting**:
   - Consolidates findings from DNS mapping, fuzzing, and port scanning into a structured final report.

---

## **How to Run**
### **1. Prerequisites**
- Install [Docker](https://www.docker.com/).
- Clone this repository:
  ```bash
  git clone https://github.com/YamatoGr93/JuiceShopRecon.git
  cd JuiceShopRecon
  ```

### **2. Build the Docker Image**
```bash
docker build -t recon_tool .
```

### **3. Run the Tool**
```bash
docker run --rm -v $(pwd)/reports:/app recon_tool http://localhost:3000
```
- Replace `http://localhost:3000` with your target URL.
- Reports will be saved in the `reports` directory.

---

## **Reports**
### **DNS Report**
The DNS mapping tool successfully retrieved records for the target `localhost`:
- **A Records**:
    - `::1`
    - `127.0.0.1`
- **CNAME Record**:
    - `localhost.home.`
- **MX and TXT Records**:
    - None (expected for `localhost`).

---

### **Fuzzing Report**
The fuzzing tool tested several payloads, uncovering potential vulnerabilities:
- **Path Traversal**:
    - `../../../etc/passwd`: Returned **200 OK**.
    - `../../../../../../etc/shadow`: Returned **200 OK**.
- **Cross-Site Scripting (XSS)**:
    - `/<script>alert(1)</script>`: Returned **200 OK**.
    - `/?q=<svg/onload=alert(1)>`: Returned **200 OK**.

**Note**: These results suggest the application may not be validating or sanitizing input properly, potentially exposing sensitive data and enabling malicious scripts.

---

### **Port Scan Report**
The port scanner identified the following:
- **Open Ports**:
    - Port `3000`: Open, service unknown (likely HTTP as Juice Shop runs here).
- **Closed Ports**:
    - Remaining ports from `3001` to `3100`.

---

## **Potential Vulnerabilities Identified**
1. **Path Traversal**:
    - Exploiting this could allow attackers to access sensitive files like `/etc/passwd` or `/etc/shadow`.

2. **Cross-Site Scripting (XSS)**:
    - Malicious scripts could be executed in users' browsers, leading to data theft or session hijacking.

3. **Open Port**:
    - Port `3000` is open and potentially exposes a web service. Further analysis is required to confirm its security posture.

---


## **About the Project**
This project demonstrates my ability to build and deploy custom security tools, leveraging Docker for portability and scalability. It is a practical application of my penetration testing and cybersecurity knowledge, aimed at identifying vulnerabilities in web applications.

If you have any questions or suggestions, feel free to reach out!

---

## **Contact**
**Panagiotis Lampros**  
Email: yamatoGr93@proton.me  
LinkedIn: [Panagiotis Lampros](https://www.linkedin.com/in/panagiotis-lampros-7b6716331/)  
GitHub: [YamatoGr93](https://github.com/YamatoGr93)
```

