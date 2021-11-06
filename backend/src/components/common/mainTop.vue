<template>
    <div class="mainTop">
        <span>{{title}}</span>
    </div>
</template>

<script>
    export default {
        name:"mainTop",
        props:{
            title:{
                type:String,
                default:"内容标题"
            }
        }
    }
</script>

<style scoped>
.mainTop {
    display: flex;
    align-items: center;
    height: 35px;
    border-bottom: solid 2px #f4f4f4;
    font-size: 16px;
}
</style>