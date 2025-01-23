# deploy arbiter

## deploy by script
1. Log in to the server

2. Enter the home directory
   ```shell
   cd ~
   ```

3. Download deploy script
   ```shell
   wget https://download.bel2.org/loan-arbiter/deploy_loan_arbiter.sh
   ```

4. Script permission changes
   ```shell
   chmod a+x deploy_loan_arbiter.sh
   ```

5. Execute deploy script
   ```shell
   ./deploy_loan_arbiter.sh [your_arbiter_esc_address] [hex_encoded_btc_private_key] [hex_encoded_esc_private_key] [your_keystore_password]
   ```
   replace  ***[your_arbiter_esc_address]*** with your esc arbiter address, not operator address, <span style="color: yellow;">with "0x"</span> at the begining.

   replace  ***[hex_encoded_btc_private_key]*** with your own arbiter btc private key, <span style="color: yellow;">without "0x"</span> at the begining.

   replace  ***[hex_encoded_esc_private_key]*** with your own esc operator private key, <span style="color: yellow;">without "0x"</span> at the begining. 

   replace  ***[your_keystore_password]*** with your own keystore password.

   For example:
   ```shell
   ./deploy_loan_arbiter.sh 0x0262aB0ED65373cC855C34529fDdeAa0e686D913 0123456789abcdef015522dd7fee2104750cb5c0be9d06d42348cf9b65c253cb0 0123456789abcdef015522dd7fee2104750cb5c0be9d06d42348cf9b65c253cb0 mypassword
   ```

   <span style="color: green;">esc private key used to submit arbiter signature to esc contract, need to have enough ESC ELA!</span>

   <span style="color: red;">The wallet cannot be used for any other purposes!</span>

   <span style="color: red;">Do not disclose your private key to avoid asset loss!</span>

6. Check arbiter status 

   check deploy script status succeed or failed.

7. Logs

   event log: ~/loan_arbiter/data/logs/event.log

   detailed arbiter log: ~/loan_arbiter/data/logs/arbiter.log

## kill arbiter

   ```shell
   pkill -x "arbiter"
   ```

## restart arbiter

   ```shell
   cd ~/loan_arbiter
   ./arbiter --gf.gcfg.file=config.yaml  > ~/loan_arbiter/data/logs/arbiter.log 2>&1 &
   ```

## deploy by docker

1. install docker

2. Enter the home directory
   ```shell
   cd ~
   ```

3. Download deploy script 
   ```shell
   wget https://download.bel2.org/loan-arbiter/docker_run_arbiter_signer.sh
   ```

4. Script permission changes
   ```shell
   chmod a+x docker_run_arbiter_signer.sh
   ```

5. Execute deploy script

   ```shell
   ./docker_run_arbiter_signer.sh
   ```