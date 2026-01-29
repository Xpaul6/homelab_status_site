import express from 'express';
import swaggerUi from 'swagger-ui-express';
import swaggerAutogen from 'swagger-autogen';
import fs from 'fs';

import si from 'systeminformation';

const app = express();
const PORT = 8080;

// Swagger
const swaggerOptions = {
  info: {
    title: 'Dashboard API',
    description: 'Homelab dashboard API',
  },
  host: 'localhost:8080'
};
const outputFile = './swagger-output.json';
const routes = ['./index.js'];

swaggerAutogen()(outputFile, routes, swaggerOptions);
const swaggerFile = JSON.parse(fs.readFileSync('./swagger-output.json'));
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerFile));

// Routes
app.get('/server-status', (_, res) => { return res.send('Server is up!') });

app.get('/mcstatus/:address', async (req, res) => {
  try {
    const apiResponse = await fetch(`https://api.mcstatus.io/v2/status/java/${req.params.address}`);
    const info = await apiResponse.json();
    return res.status(200).send(info);
  } catch (err) {
    console.log(err);
    return res.status(500).send({ error: 'Failed to fetch Minecraft server status' });
  }
});

app.get('/websitestatus', async (req, res) => {
  try {
    const siteResponse = await fetch(req.query.address);
    if (siteResponse.status == 200) {
      return res.status(200).send();
    } else {
      return res.status(503).send();
    }
  } catch (err) {
    console.log(err);
    return res.status(500).send({ error: err });
  }
});

app.get('/sysinfo', async (_, res) => {
  try {
    const [temp, load, mem] = await Promise.all([
      si.cpuTemperature(),
      si.currentLoad(),
      si.mem()
    ]);

    const result = {
      cpu_temp: temp.main,
      cpu_load_percent: load.currentLoad,
      ram_percent: mem.active / mem.total * 100
    };

    return res.status(200).send(result);
  } catch (error) {
    console.log(error);
    return res.status(500).send({ error: 'Failed to fetch system information' });
  }
});

// Server
app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`);
});
