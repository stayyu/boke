<template >
    <div>
        <Toolbar :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" style="border-bottom: 1px solid #ccc" />
        <Editor :defaultConfig="editorConfig" :mode="mode" v-model="valueHtml" style="height: 400px; overflow-y: hidden"
            @onCreated="handleCreated" @onChange="handleChange" />
    </div>
</template>

<script setup>
import { ref, reactive, inject, onMounted, onBeforeUnmount, shallowRef } from 'vue'
import '@wangeditor/editor/dist/css/style.css';
import { Adminstore } from '../stores/Adminstore';
import { Editor, Toolbar } from '@wangeditor/editor-for-vue';

const server_url = inject("server_url")
// 编辑器实例，必须用 shallowRef，重要！
const editorRef = shallowRef()
const mode = ref("default")
const toolbarConfig={excludeKeys:["uploadVideo"]}
const editorConfig = { placeholder: '请输入内容...' }
const valueHtml = ref()
editorConfig.MENU_CONF = {}
editorConfig.MENU_CONF['uploadImage'] = {
    server: server_url + '/bk/bok/upload',
}
editorConfig.MENU_CONF['insertImage'] = {
 onSuccess(file, res) {          // JS 语法
        console.log(`${file.name} 上传成功`, res)
    }, onError(file, err, res) {               // JS 语法
        console.log(`${file.name} 上传出错`, err, res)
    },
 
}


const props = defineProps({
    modelValue: {
        modelValue: {
            type: String,
            default: ""
        }
    }
})

const emit = defineEmits(["update:model-value"])

let initFinshed = false

onMounted(() => {
    setTimeout(() => {
        valueHtml.value = props.modelValue;
        initFinshed = true;
    }, 10)
})



// 组件销毁时，也及时销毁编辑器，重要！
onBeforeUnmount(() => {
    const editor = editorRef.value;
    if (editor == null) return;

    editor.destroy();
});

// 编辑器回调函数
const handleCreated = (editor) => {
    console.log('created', editor);
    editorRef.value = editor; // 记录 editor 实例，重要！
};
const handleChange = (editor) => {
    emit("update:model-value", valueHtml.value)//提交原始值
};
</script>
<style lang="scss" scoped></style>