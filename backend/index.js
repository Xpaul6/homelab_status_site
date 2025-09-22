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

// Server
app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`)
})
