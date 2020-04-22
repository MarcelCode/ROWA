<template>
    <v-row align-center>
        <v-col cols="2" style="padding-left: 30px">
            <v-btn dark fab color="white" :to="{name:this.prevPage}">
                <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
        </v-col>
        <v-col align="center"  align-self="center" cols="8">
             <h3 style="color:#828282" v-text="headtext"></h3>
        </v-col>
        <v-col cols="2" id="logo-text" style="text-align: right; padding-right: 30px">
            <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
        </v-col>
    </v-row>

</template>

<script>
    
    export default {
        name: "AdminTopRow",
        components:{
            
        },
        props:{
                headtext: String,
                prevPage: String
        }
    }
</script>

<style scoped>
    #logo-text{
        font-family: Montserrat, sans-serif;
        font-style: normal;
        font-weight: 600;
        font-size: 36px;
        color: var(--v-primary-base)
    }
</style>