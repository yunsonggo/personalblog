<template>
    <div>
        <div class="personal_edit" @touchstart="editClick">
            <span>{{left}}</span>
            <span><slot name="right" /></span>
        </div>
    </div>
</template>

<script>
    export default {
        name:"personalEdit",
        props:{
            left:{
                type:String,
                default:""
            },
            isInput:{
                type:Boolean,
                default:true
            },
            rule:{
                type:String,
                default:""
            }
        },
        methods:{
            editClick() {
                this.$emit("editClick",this.left,this.rule)
            }
        }
    }
</script>

<style scoped>
.personal_edit {
    background-color: white;
    padding:3.333vw;
    border-bottom: 0.278vw solid #eee;
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.personal_edit span{
    color:#333333;
    font-size: 14px;
}
.personal_edit span:nth-child(1){
    width: 12.5vw;
}
</style>