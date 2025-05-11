import * as d3 from 'd3'

import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { SpaghettiCannonApiService } from "./ts/proto/SpaghettiCannon/clientapi/v1/clientapi_pb"

import { createApp } from 'vue'
import MagicButton from './vue/MagicButton.vue'
import UpdateBox from './vue/UpdateBox.vue'

function main (): void {
  document.querySelector('#history').appendChild(createGraph())

  /**
  createApp(MagicButton).mount('#magicButton')

  createApp(UpdateBox).mount('#updateBox')
  */

  let baseUrl = window.location.href + 'api/'

  if (window.location.port == '5173') {
	  baseUrl = 'http://localhost:4337/api/'
  }
  const transport = createConnectTransport({
    baseUrl: baseUrl
  })

  window.client = createClient(SpaghettiCannonApiService, transport)

  checkReady()
}

async function checkReady (): void {
  const res = await window.client.getReadyz({
    name: "Batman",
  })

  document.getElementById('version').innerText = res.version

  const icons = {
		"weight-outline": "Weight",
		"heart-outline": "Heart",
		"mental-disorders-outline": "Mood",
		"exercise-bicycle-outline": "Exercise",
		"blood-drop-outline": "BP",
		"unhealthy-food-outline": "Food",
  }

  for (const [icon, title] of Object.entries(icons)) {
	  let li = document.createElement('li')

	  let a = document.createElement('a')
	  a.classList.add('icon');

	  let domIcon = document.createElement('iconify-icon')
	  domIcon.setAttribute('icon', 'healthicons:' + icon)
	  a.appendChild(domIcon)

	  li.appendChild(a)

	  document.querySelector('nav').appendChild(li)
  }

  console.log("readyz", res);
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
