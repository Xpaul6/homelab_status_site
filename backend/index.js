import express from 'express'
import swaggerUi from 'swagger-ui-express'
import swaggerJsDoc from 'swagger-jsdoc'

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
app.get('/server-status', (_, res) => { res.send('Server is up!') })

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
    res.status(200).send(info);
  } catch (err) {
    console.log(err);
    res.status(500).send({ error: 'Failed to fetch Minecraft server status' });
  }
})

// Server
app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`)
})
