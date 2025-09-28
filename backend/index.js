import express from 'express'
import swaggerUi from 'swagger-ui-express'
import swaggerJsDoc from 'swagger-jsdoc'

import si from 'systeminformation'

const app = express()
const PORT = 8080

// Swagger
const swaggerOptions = {
  swaggerDefinition: {
    openapi: '3.0.0',
    info: {
      title: 'My API',
      version: '1.0.0',
      description: 'API documentation',
    },
    servers: [
      {
        url: 'http://localhost:8080',
      },
    ],
  },
  apis: ['./index.js'],
}
const swaggerDocs = swaggerJsDoc(swaggerOptions);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocs));

// Routes
/**
 * @swagger
 * /server-status:
 *   get:
 *     description: Returns server status message string
 *     responses:
 *       200:
 *         description: A successful response.
 */
app.get('/server-status', (_, res) => { return res.send('Server is up!') })

/**
 * @swagger
 * /mcstatus/{address}:
 *   get:
 *     summary: Gets the status of a Minecraft server.
 *     description: Fetches the status of a given Minecraft server address from the mcstatus.io API.
 *     parameters:
 *       - in: path
 *         name: address
 *         required: true
 *         description: The address (IP or domain) of the Minecraft server.
 *         schema:
 *           type: string
 *     responses:
 *       200:
 *         description: Successful response containing the server status.
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               example:
 *                 online: true
 *                 host: "demo.mcstatus.io"
 *                 port: 25565
 *                 players:
 *                   online: 0
 *                   max: 20
 *                 version:
 *                   name_raw: "1.19.4"
 *                   protocol: 762
 *       500:
 *         description: Error fetching the server status.
 */
app.get('/mcstatus/:address', async (req, res) => {
  try {
    const apiResponse = await fetch(`https://api.mcstatus.io/v2/status/java/${req.params.address}`);
    const info = await apiResponse.json();
    return res.status(200).send(info);
  } catch (err) {
    console.log(err);
    return res.status(500).send({ error: 'Failed to fetch Minecraft server status' });
  }
})

/**
 * @swagger
 * /sysinfo:
 *   get:
 *     summary: Gets system information.
 *     description: Returns basic system information like CPU temperature, CPU load, and RAM usage.
 *     responses:
 *       200:
 *         description: Successful response containing the system information.
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 cpu_temp:
 *                   type: number
 *                   description: "CPU temperature in Celsius."
 *                   example: 45.5
 *                 cpu_load:
 *                   type: number
 *                   description: "CPU load in percent."
 *                   example: 15.2
 *                 ram_percent:
 *                   type: number
 *                   description: "RAM usage as a percentage."
 *                   example: 0.65
 *       500:
 *         description: Error fetching the system information.
 */
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
})

// Server
app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`)
})
