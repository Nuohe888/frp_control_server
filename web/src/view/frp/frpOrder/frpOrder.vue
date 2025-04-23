<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" :rules="searchRule" class="demo-form-inline"
               @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="CreatedAt">
          <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt"
                          :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"
                          placeholder="开始日期"
                          type="datetime"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt"
                          :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"
                          placeholder="结束日期"
                          type="datetime"></el-date-picker>
        </el-form-item>


        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button icon="search" type="primary" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button v-if="!showAllQuery" icon="arrow-down" link type="primary" @click="showAllQuery=true">展开
          </el-button>
          <el-button v-else icon="arrow-up" link type="primary" @click="showAllQuery=false">收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="plus" type="primary" @click="openDialog()">新增</el-button>
        <el-button :disabled="!multipleSelection.length" icon="delete" style="margin-left: 10px;" @click="onDelete">
          删除
        </el-button>

      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        row-key="ID"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="CreatedAt" sortable width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="内部UUID" prop="uuid" width="120" />

        <el-table-column align="left" label="限速" prop="speed" width="120" />

        <el-table-column align="left" label="流量" prop="flow" width="120" />
        <el-table-column align="left" label="端口" prop="port" width="120" />

        <el-table-column :min-width="appStore.operateMinWith" align="left" fixed="right" label="操作">
          <template #default="scope">
            <el-button class="table-button" link type="primary" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看
            </el-button>
            <el-button class="table-button" icon="edit" link type="primary" @click="updateFrpOrderFunc(scope.row)">
              编辑
            </el-button>
            <el-button icon="delete" link type="primary" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer v-model="dialogFormVisible" :before-close="closeDialog" :show-close="false" :size="appStore.drawerSize"
               destroy-on-close>
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form ref="elFormRef" :model="formData" :rules="rule" label-position="top" label-width="80px">
        <el-form-item label="内部UUID:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="false" placeholder="请输入内部UUID" />
        </el-form-item>
        <el-form-item label="限速:" prop="speed">
          <el-input v-model.number="formData.speed" :clearable="false" placeholder="请输入限速" />
        </el-form-item>
        <el-form-item label="流量:" prop="flow">
          <el-input v-model.number="formData.flow" :clearable="false" placeholder="请输入流量" />
        </el-form-item>
        <el-form-item label="节点ID:" prop="nodeId">
          <el-select v-model="formData.nodeId" :clearable="false" filterable placeholder="请选择节点ID"
                     style="width:100%">
            <el-option v-for="(item,key) in dataSource.nodeId" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="端口:" prop="port">
          <el-input v-model.number="formData.port" :clearable="true" placeholder="请输入端口" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="detailShow" :before-close="closeDetailShow" :show-close="true" :size="appStore.drawerSize"
               destroy-on-close title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="内部UUID">
          {{ detailFrom.uuid }}
        </el-descriptions-item>
        <el-descriptions-item label="限速">
          {{ detailFrom.speed }}
        </el-descriptions-item>
        <el-descriptions-item label="流量">
          {{ detailFrom.flow }}
        </el-descriptions-item>
        <el-descriptions-item label="端口">
          {{ detailFrom.port }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    getFrpOrderDataSource,
    createFrpOrder,
    deleteFrpOrder,
    deleteFrpOrderByIds,
    updateFrpOrder,
    findFrpOrder,
    getFrpOrderList
  } from '@/api/frp/frpOrder'

  // 全量引入格式化工具 请按需保留
  import {
    getDictFunc,
    formatDate,
    formatBoolean,
    filterDict,
    filterDataSource,
    returnArrImg,
    onDownloadFile
  } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'
  import { useAppStore } from '@/pinia'


  defineOptions({
    name: 'FrpOrder'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    uuid: '',
    speed: undefined,
    flow: undefined,
    nodeId: undefined,
    port: undefined
  })
  const dataSource = ref([])
  const getDataSourceFunc = async () => {
    const res = await getFrpOrderDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()


  // 验证规则
  const rule = reactive({
    uuid: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
      }
    ],
    speed: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }
    ],
    flow: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }
    ],
    nodeId: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }
    ],
    port: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    }
    ]
  })

  const searchRule = reactive({
    CreatedAt: [
      {
        validator: (rule, value, callback) => {
          if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
            callback(new Error('请填写结束日期'))
          } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
            callback(new Error('请填写开始日期'))
          } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
            callback(new Error('开始日期应当早于结束日期'))
          } else {
            callback()
          }
        }, trigger: 'change'
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      page.value = 1
      getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getFrpOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {
  }

  // 获取需要的字典 可能为空 按需保留
  setOptions()


  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteFrpOrderFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
      const res = await deleteFrpOrderByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateFrpOrderFunc = async (row) => {
    const res = await findFrpOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteFrpOrderFunc = async (row) => {
    const res = await deleteFrpOrder({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      uuid: '',
      speed: undefined,
      flow: undefined,
      nodeId: undefined
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    btnLoading.value = true
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return btnLoading.value = false
      let res
      switch (type.value) {
        case 'create':
          res = await createFrpOrder(formData.value)
          break
        case 'update':
          res = await updateFrpOrder(formData.value)
          break
        default:
          res = await createFrpOrder(formData.value)
          break
      }
      btnLoading.value = false
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      }
    })
  }

  const detailFrom = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)


  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }


  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findFrpOrder({ ID: row.ID })
    if (res.code === 0) {
      detailFrom.value = res.data
      openDetailShow()
    }
  }


  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailFrom.value = {}
  }


</script>

<style>

</style>
