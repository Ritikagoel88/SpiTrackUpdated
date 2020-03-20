var express = require("express");
var router = express.Router();
var bodyParser = require("body-parser");

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

const txSubmit = require("./invoke");
const txFetch = require("./query");

//var TFBC = require("./FabricHelper");

// Request PO
router.post("/requestPO", async function(req, res) {
  try {
    let result = await txSubmit.invoke("requestPO", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Accept PO
router.post("/acceptPO", async function(req, res) {
  try {
    let result = await txSubmit.invoke("acceptPO", (req.body.purchaseOrderId));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Reject PO
router.post("/rejectPO", async function(req, res) {
  try {
    let result = await txSubmit.invoke("rejectPO", (req.body.purchaseOrderId));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Get PO Details
router.post("/getPODetails", async function(req, res) {
  //TFBC.getLC(req, res); req.body.lcId
  try {
    let result = await txFetch.query("getPODetails", req.body.purchaseOrderId);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  } 
});

// Get PO History Details
  router.post("/getPOHistory", async function(req, res) {
    try {
      let result = await txFetch.query("getPOHistory", req.body.purchaseOrderId);
      res.send(JSON.parse(result));
    } catch (err) {
      res.status(500).send(err);
    }
  });

module.exports = router;
