import * as d3 from 'd3'

import 'femtocrank/style.css'

import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { SpaghettiCannonApiService } from "./ts/proto/SpaghettiCannon/clientapi/v1/clientapi_pb"

import { createApp } from 'vue'
import router from './router'
import App from './App.vue'
import { loadInit } from './resources/vue/composables/useInit.js'

async function main (): Promise<void> {
  const history = document.querySelector('#history')
  if (history) {
    history.appendChild(createGraph())
  }

  const transport = createConnectTransport({
    baseUrl: window.location.origin + '/api/',
  })

  window.client = createClient(SpaghettiCannonApiService, transport)

  try {
    await loadInit()
  } catch {
    // App still mounts; footer/settings use defaults until Init succeeds.
  }

  const app = createApp(App)
  app.use(router)
  app.mount('#app')
}

function createGraph (): Node {
  const width = 640
  const height = 400
  const margin = 30

  const data = [
    { date: new Date('2024-03-01'), value: 50 },
    { date: new Date('2024-03-02'), value: 52 },
    { date: new Date('2024-03-04'), value: 49 }
  ]

  const x = d3.scaleUtc().domain([new Date('2024-03-01'), new Date('2024-03-30')]).range([margin, width - margin])
  const y = d3.scaleLinear().domain([0, 100]).range([height - margin, margin])

  const svg = d3.create('svg').attr('width', width).attr('height', height)
  svg.append('g').attr('transform', `translate(0,${height - margin})`).call(d3.axisBottom(x))
  svg.append('g').attr('transform', `translate(${margin}, 0)`).call(d3.axisLeft(y))
  svg.append('path')
    .datum(data)
    .attr('fill', 'none')
    .attr('stroke', 'black')
    .attr('stroke-width', 1.5)
    .attr('d', d3.line()
      .x((d) => x(d.date))
      .y((d) => y(d.value))
    )

  return svg.node()
}

main()
