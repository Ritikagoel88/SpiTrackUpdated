/*
 * SPDX-License-Identifier: Apache-2.0
 */

"use strict";

const {
  FileSystemWallet,
  Gateway,
  X509WalletMixin
} = require("fabric-network");
const FabricCAServices = require("fabric-ca-client");
const path = require("path");
const fs = require("fs");

const ccpPath = path.resolve(__dirname, "connection-Importer.json");
const ccpJSON = fs.readFileSync(ccpPath, "utf8");
const ccp = JSON.parse(ccpJSON);
async function main() {
  try {
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), "wallet");
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.exists("ImporterUser");
    if (userExists) {
      console.log(
        'An identity for the user "ImporterUser" already exists in the wallet'
      );
      return;
    }

    // Check to see if we've already enrolled the 0 user.
    const adminExists = await wallet.exists("admin-Importer");
    if (!adminExists) {
      const caInfo = ccp.certificateAuthorities["ca.Importer.spitrack.com"];  // SP: replaced bank with spitrack
      console.log(caInfo);
      //const caTLSCACerts = caInfo.tlsCACerts.pem;
      const ca = new FabricCAServices(
        caInfo.url,
        { verify: false },
        caInfo.caName
      );
      // Enroll the admin user, and import the new identity into the wallet.
      const enrollment = await ca.enroll({
        enrollmentID: "admin",
        enrollmentSecret: "adminpw"
      });
      const identity = X509WalletMixin.createIdentity(
        "ImporterMSP",
        enrollment.certificate,
        enrollment.key.toBytes()
      );
      await wallet.import("admin-Importer", identity);
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccpPath, {
      wallet,
      identity: "admin-Importer",
      discovery: { enabled: true, asLocalhost: true }
    });

    // Get the CA client object from the gateway for interacting with the CA.
    const ca = gateway.getClient().getCertificateAuthority();
    const adminIdentity = gateway.getCurrentIdentity();

    // Register the user, enroll the user, and import the new identity into the wallet.
    const secret = await ca.register(
      { enrollmentID: "ImporterUser", role: "client" },
      adminIdentity
    );
    const enrollment = await ca.enroll({
      enrollmentID: "ImporterUser",
      enrollmentSecret: secret
    });
    const userIdentity = X509WalletMixin.createIdentity(
      "ImporterMSP",
      enrollment.certificate,
      enrollment.key.toBytes()
    );
    await wallet.import("ImporterUser", userIdentity);
    console.log(
      'Successfully registered and enrolled admin user "ImporterUser" and imported it into the wallet'
    );
  } catch (error) {
    console.error(`Failed to register user "ImporterUser": ${error}`);
    process.exit(1);
  }
}

main();