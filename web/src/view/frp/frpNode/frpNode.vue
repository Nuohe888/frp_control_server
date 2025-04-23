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
          <el-date-picker v-model="searchInfo.startCreatedAt" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false" placeholder="开始日期"
                          type="datetime"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false" placeholder="结束日期"
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

        <el-table-column align="left" label="描述" prop="desc" width="120" />

        <el-table-column align="left" label="名称标识" prop="name" width="120" />

        <el-table-column align="left" label="节点密钥" prop="key" width="320" />

        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status, frpNodeStatusOptions) }}
          </template>
        </el-table-column>
        <el-table-column :min-width="appStore.operateMinWith" align="left" fixed="right" label="操作">
          <template #default="scope">
            <el-button class="table-button" link type="primary" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看
            </el-button>
            <el-button class="table-button" icon="edit" link type="primary" @click="updateFrpNodeFunc(scope.row)">编辑
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
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入描述" />
        </el-form-item>
        <el-form-item label="名称标识:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称标识" />
        </el-form-item>
        <el-form-item label="节点密钥:" prop="key">
          <el-input v-model="formData.key" :clearable="true" placeholder="请输入节点密钥" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" :clearable="true" filterable placeholder="请选择状态" style="width:100%">
            <el-option v-for="(item,key) in frpNodeStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="detailShow" :before-close="closeDetailShow" :show-close="true" :size="appStore.drawerSize"
               destroy-on-close title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="描述">
          {{ detailFrom.desc }}
        </el-descriptions-item>
        <el-descriptions-item label="名称标识">
          {{ detailFrom.name }}
        </el-descriptions-item>
        <el-descriptions-item label="节点密钥">
          {{ detailFrom.key }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          {{ detailFrom.status }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
  import {
    createFrpNode,
    deleteFrpNode,
    deleteFrpNodeByIds,
    updateFrpNode,
    findFrpNode,
    getFrpNodeList
  } from '@/api/frp/frpNode'

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
    name: 'FrpNode'
  })

  // 提交按钮loading
  const btnLoading = ref(false)
  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const frpNodeStatusOptions = ref([])
  const formData = ref({
    desc: '',
    name: '',
    key: '',
    status: ''
  })


  // 验证规则
  const rule = reactive({
    name: [{
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
    key: [{
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
    status: [{
      required: true,
      message: '',
      trigger: ['input', 'blur']
    },
      {
        whitespace: true,
        message: '不能只输入空格',
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
    const table = await getFrpNodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    frpNodeStatusOptions.value = await getDictFunc('frpNodeStatus')
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
      deleteFrpNodeFunc(row)
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
      const res = await deleteFrpNodeByIds({ IDs })
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
  const updateFrpNodeFunc = async (row) => {
    const res = await findFrpNode({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 删除行
  const deleteFrpNodeFunc = async (row) => {
    const res = await deleteFrpNode({ ID: row.ID })
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
      desc: '',
      name: '',
      key: '',
      status: ''
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
          res = await createFrpNode(formData.value)
          break
        case 'update':
          res = await updateFrpNode(formData.value)
          break
        default:
          res = await createFrpNode(formData.value)
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
    const res = await findFrpNode({ ID: row.ID })
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
