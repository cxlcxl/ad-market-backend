<template>
  <dialog-panel title="添加课程" confirm-text="保存" :visible="visible" @cancel="cancel" @confirm="save"
                :confirm-loading="loading" width="800px">
    <el-form :model="listenForm" ref="listenForm" label-width="100px" size="small" :rules="listenRules">
      <el-row :gutter="30">
        <el-col :span="12">
          <el-form-item label="课程图" prop="img_url">
            <el-input placeholder="请选择课程图" v-model="listenForm.img_url" disabled>
              <el-button slot="append" icon="el-icon-thumb" @click="selectImg">选择</el-button>
            </el-input>
          </el-form-item>
          <el-form-item label="标题" prop="title">
            <el-input v-model="listenForm.title" placeholder="请填写课程标题" />
          </el-form-item>
          <el-form-item label="小标题" prop="sub_title">
            <el-input v-model="listenForm.sub_title" placeholder="请填写课程小标题" />
          </el-form-item>
          <el-form-item label="排序" prop="order_by">
            <el-input-number v-model="listenForm.order_by" :min="1" :max="9999"/>
          </el-form-item>
        </el-col>
        <el-col :span="12" class="list-box">
          <el-row :gutter="10" v-for="(ls, idx) in listenForm.lists">
            <el-col :span="16">
              <el-form-item :prop="'lists.'+idx+'.title'"
                            label-width="0"
                            :rules="{required: true, message: '请填写目录标题'}">
                <el-input v-model="ls.title" placeholder="目录标题(必填)"/>
              </el-form-item>
            </el-col>
            <el-col :span="5">
              <el-form-item :prop="'lists.'+idx+'.order_by'" label-width="0" :rules="{required: true, message: '请填写排序'}">
                <el-input-number v-model="ls.order_by" :min="1" :max="9999" :controls="false" style="width: 100%;"/>
              </el-form-item>
            </el-col>
            <el-col :span="3">
              <el-form-item label-width="0" v-show="idx>0">
                <el-button type="danger" icon="el-icon-delete" circle size="mini" @click="removeOne(idx)"></el-button>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24" style="text-align: center;">
              <el-form-item label-width="0">
                <el-button type="success" plain icon="el-icon-plus" @click="moreList">追加一条目录</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-form>

    <listen-img-select ref="img_select" @handle-select="selectedImg"/>
  </dialog-panel>
</template>

<script>
  import DialogPanel from "@c/DialogPanel"
  import { listenCreate } from "@a/listen"
  import ListenImgSelect from "./select-images"

  export default {
    components: {
      DialogPanel, ListenImgSelect
    },
    data() {
      return {
        visible: false,
        loading: false,
        listenForm: {
          id: 0,
          title: "",
          img_url: "",
          sub_title: "",
          order_by: 99,
          lists: [{title: "", order_by: 1}],
        },
        listenRules: {
          img_url: { required: true, message: "请选择课程图" },
          title: { required: true, message: "请填写标题" },
          sub_title: { required: true, message: "请填写小标题" },
        },
      }
    },
    methods: {
      initCreate() {
        this.visible = true
      },
      cancel() {
        this.$refs.listenForm.resetFields()
        this.visible = false
      },
      moreList() {
        let orderBy = this.listenForm.lists[this.listenForm.lists.length-1].order_by+1
        this.listenForm.lists.push({title: "", order_by: orderBy})
      },
      removeOne(idx) {
        this.listenForm.lists.splice(idx, 1)
      },
      selectImg() {
        this.$refs.img_select.initSelect()
      },
      selectedImg(v) {
        this.$set(this.listenForm, "img_url", v)
      },
      save() {
        this.$refs.listenForm.validate((v) => {
          if (v) {
            this.loading = true
            listenCreate(this.listenForm)
              .then((res) => {
                this.$message.success("添加成功")
                this.$emit("success")
                this.loading = false
                this.cancel()
              })
              .catch((err) => {
                this.loading = false
              })
          } else {
            return false
          }
        })
      },
    },
  }
</script>

<style lang="less">
  .list-box {
    border-left: 1px solid #DCDFE6;
    max-height: 400px;
    overflow-y: scroll;

  &::-webkit-scrollbar {
     /*滚动条整体样式*/
     width: 6px; /*高宽分别对应横竖滚动条的尺寸*/
     height: 0;
   }

  &::-webkit-scrollbar-thumb {
     /*滚动条里面小方块*/
     border-radius: 10px;
     background-color: rgb(172, 184, 206);
   }

  &::-webkit-scrollbar-track {
     /*滚动条里面轨道*/
     box-shadow: inset 0 0 5px rgba(75, 73, 73, 0.2);
     background: #ededed;
     border-radius: 2px;
   }
  }
</style>
