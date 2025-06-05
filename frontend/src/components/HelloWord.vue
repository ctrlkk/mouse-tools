<script setup lang="ts">
import * as echarts from 'echarts/core'
import {
    LineChart,
    type LineSeriesOption
} from 'echarts/charts'
import {
    GridComponent,
    TooltipComponent,
    type GridComponentOption,
    type TooltipComponentOption
} from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import dayjs from 'dayjs'

// @ts-ignore
import { EventsOn } from '../../wailsjs/runtime/runtime';

echarts.use([
    GridComponent,
    TooltipComponent,
    LineChart,
    CanvasRenderer
])

type EChartsOption = echarts.ComposeOption<
    | LineSeriesOption
    | GridComponentOption
    | TooltipComponentOption
>

// 接口定义
interface PollingRateRecord {
    // 回报率值
    rate: number
    // 记录时间
    timestamp: number
}

// 组件常量配置
const CHART_CONFIG = {
    // 最大数据点数量
    MAX_DATA_POINTS: 100,
} as const

// 响应式状态
const chartContainerRef = ref<HTMLDivElement>()
const isMonitoring = ref<boolean>(false)

let chartInstance: echarts.ECharts

// 图表数据状态
const chartDataPoints = ref<(number | null)[]>(
    new Array(CHART_CONFIG.MAX_DATA_POINTS).fill(null)
)
const currentPollingRate = ref<number>(0)

// 回报率记录
const allPollingRates = ref<PollingRateRecord[]>([])

const initializeChart = () => {
    chartInstance = echarts.init(chartContainerRef.value)

    const chartOption: EChartsOption = {
        grid: {
            left: '3%',
            right: '4%',
            bottom: '8%',
            top: '5%',
            containLabel: true
        },
        xAxis: {
            type: 'category',
            data: Array.from({ length: CHART_CONFIG.MAX_DATA_POINTS }, (_, index) => index),
            axisLabel: {
                show: true,
                interval: 19,
                color: '#6b7280',
                fontSize: 11
            },
            axisLine: {
                lineStyle: {
                    color: '#e5e7eb'
                }
            },
            axisTick: {
                show: false
            }
        },
        yAxis: {
            type: 'value',
            min: 0,
            axisLabel: {
                formatter: '{value}Hz',
                color: '#6b7280',
                fontSize: 11
            },
            axisLine: {
                lineStyle: {
                    color: '#e5e7eb'
                }
            },
            splitLine: {
                lineStyle: {
                    color: '#f3f4f6',
                    type: 'dashed'
                }
            }
        },
        tooltip: {
            trigger: 'axis',
            backgroundColor: 'rgba(0, 0, 0, 0.8)',
            borderColor: 'transparent',
            textStyle: {
                color: '#ffffff',
                fontSize: 12
            },
            formatter: (params: any) => {
                const dataPoint = params[0]
                if (!dataPoint.value) return ''

                return `
                        <div style="padding: 4px 0;">
                        <div style="margin-bottom: 4px;">
                            <span style="color: #10b981;">●</span> 
                            <span style="margin-left: 4px;">${dataPoint.value}Hz</span>
                        </div>
                        <div style="font-size: 11px; color: #d1d5db;">
                            ${dayjs().format('HH:mm:ss')}
                        </div>
                        </div>
                    `
            }
        },
        series: [{
            name: '鼠标回报率',
            type: 'line',
            data: chartDataPoints.value,
            smooth: false,
            symbol: 'circle',
            symbolSize: 3,
            lineStyle: {
                color: '#10b981',
                width: 2
            },
            itemStyle: {
                color: '#10b981'
            },
            areaStyle: {
                color: {
                    type: 'linear',
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    colorStops: [
                        { offset: 0, color: 'rgba(16, 185, 129, 0.3)' },
                        { offset: 1, color: 'rgba(16, 185, 129, 0.05)' }
                    ]
                }
            },
            connectNulls: false,
            animation: false
        }]
    }

    chartInstance.setOption(chartOption)

    window.addEventListener('resize', handleWindowResize)
}

/**
 * 更新图表数据
 * @param newPollingRate 新的回报率值
 */
const updateChartData = (newPollingRate: number): void => {
    for (let i = 0; i < CHART_CONFIG.MAX_DATA_POINTS - 1; i++) {
        chartDataPoints.value[i] = chartDataPoints.value[i + 1]
    }
    chartDataPoints.value[CHART_CONFIG.MAX_DATA_POINTS - 1] = newPollingRate

    if (chartInstance) {
        chartInstance.setOption({
            series: [{
                data: chartDataPoints.value
            }]
        }, false)
    }

    currentPollingRate.value = newPollingRate

    allPollingRates.value.push({
        rate: newPollingRate,
        timestamp: Date.now()
    })

    if (allPollingRates.value.length > 1000) {
        allPollingRates.value = allPollingRates.value.slice(-500)
    }
}

let stopMouseRateListener = true

EventsOn("MouseRate", (count: number) => {
    if (stopMouseRateListener) return
    updateChartData(count)
})

/**
 * 切换监控状态
 */
const handleToggleMonitoring = (): void => {
    isMonitoring.value = !isMonitoring.value

    if (isMonitoring.value) {
        stopMouseRateListener = false;
    } else {
        stopMouseRateListener = true;
    }
}

/**
 * 清空所有数据
 */
const handleClearData = (): void => {
    chartDataPoints.value = new Array(CHART_CONFIG.MAX_DATA_POINTS).fill(null)
    currentPollingRate.value = 0
    allPollingRates.value = []

    if (chartInstance) {
        chartInstance.setOption({
            series: [{
                data: chartDataPoints.value
            }]
        }, false)
    }
}

/**
 * 处理窗口大小变化
 */
const handleWindowResize = (): void => {
    if (chartInstance) {
        chartInstance.resize()
    }
}

onMounted(() => {
    initializeChart()
    handleToggleMonitoring()
})

onUnmounted(() => {
    window.removeEventListener('resize', handleWindowResize)
    chartInstance?.dispose()
})
</script>

<template>
    <div class="w-full">
        <NCard>
            <template #header>
                <NButtonGroup size="small">
                    <NButton @click="handleToggleMonitoring" :type="isMonitoring ? 'warning' : 'success'" ghost>
                        {{ isMonitoring ? '停止监控' : '开始监控' }}
                    </NButton>

                    <NButton @click="handleClearData" type="error" ghost>
                        清空数据
                    </NButton>
                </NButtonGroup>
            </template>
            <template #default>
                <div ref="chartContainerRef" class="h-80 p-3"></div>
            </template>
        </NCard>
    </div>
</template>

<style scoped></style>
