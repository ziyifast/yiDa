export const mixin = {
    methods: {
        //提示信息
        notify(titile, type) {
            this.$notify({
                title: titile,
                type: type,
            })
        }
    }
}