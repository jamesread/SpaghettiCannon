import * as d3 from 'd3'

function main (): void {
  document.querySelector('section.history').querySelector('.pagewidth').appendChild(createGraph())
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
