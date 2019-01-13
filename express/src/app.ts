const express = require("express");
const bodyParser = require("body-parser");
const app = express();
const port = 3000;
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.get("/", (req, res) => res.send("Hello World!"));

app.post("/", (req, res) => {
  let text = req.body.text;
  console.log(text);
  let data = {
    response_type: "in_channel",
    text: "こんにちは、<@" + req.body.user_id + ">さん"
  };
  res.json(data);
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
