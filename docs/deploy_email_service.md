# email service

## deploy
1. Log in to the server, and copy **arbiter_email** to **"~/loan_arbiter"**

2. Enter the work directory
   ```shell
   cd ~/loan_arbiter
   ```

3. Prepare config.yaml

   ```yaml
   chain:
     esc: "https://api.elastos.io/esc"

   arbiter:
     escStartHeight: 28450960
     loanNeedSignReqPath: "~/loan_arbiter/email/data/request"
     escArbiterContractAddress: "0xA10b92006743Ef3B12077da67e465963743b03D3"
     arbiters: ["0xa20f5A22eF423b0e5c2Db79A5475D9512d989971", "0x0262aB0ED65373cC855C34529fDdeAa0e686D913"]

   email:
     host: "your_host"
     port: 587
     from: "your_email"
     user: "your_username"
     password: "your_password"
     to: ["to_email_1", "to_email_2"]
     emailLogPath: "~/loan_arbiter/email/data/logs"
     dataPath: "~/loan_arbiter/email/data"
     duration: 3600 # seconds
   ```

4. Execute the arbiter_email

   ```shell
   ./arbiter_email --gf.gcfg.file=config.yaml  > ~/loan_arbiter/email/data/logs/email.log 2>&1 &
   ```
   