<template>
  <div ref="chartRef" style="height: 300px; border: 1px solid red;"></div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts/core'
import { LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent } from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'

echarts.use([
  GridComponent,
  TooltipComponent,
  LineChart,
  CanvasRenderer
])

const chartRef = ref<HTMLDivElement>()

const tooltipFormatter = (params: any) => {
  console.log('Tooltip formatter triggered!', params)
  return 'Test Tooltip'
}

onMounted(() => {
  const chart = echarts.init(chartRef.value!)
  chart.setOption({
    tooltip: {
      trigger: 'axis',
      formatter: tooltipFormatter,
      showContent: true
    },
    xAxis: {
      type: 'category',
      data: ['A', 'B', 'C', 'D']
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        type: 'line',
        data: [10, 20, 30, 40]
      }
    ]
  })
})
</script>
