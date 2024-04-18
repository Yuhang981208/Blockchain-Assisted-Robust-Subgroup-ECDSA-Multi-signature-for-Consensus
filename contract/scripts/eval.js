const fs = require("fs");
const Registry = artifacts.require("Registry");

module.exports = async function () {

  let registry = await Registry.deployed();
  let message = "0x930d2dedab40cb4c03a967aea4f54b22ba6328f7096dc44590e651de6e2a416b";


  await registry.requestSign(message);

};
