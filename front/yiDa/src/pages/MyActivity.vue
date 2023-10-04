<template>
    <div>
        <h1>{{ username }} 的 个人中心</h1>
    </div>
</template>

<script>
export default {
    data() {
        return {

        }
    },
    computed: {
        username() {
            return sessionStorage.getItem('username')
        }
    }
}

</script>