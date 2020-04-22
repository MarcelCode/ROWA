<template>
    <div style="position: relative; width: 410px; height: 522px" >
        <svg width="410" height="522" viewBox="0 0 410 522" fill="none" xmlns="http://www.w3.org/2000/svg">
            <rect x="190" width="30" height="494" fill="#828282"/>
            <path d="M0 522L0 494H410V522H0Z" fill="#828282"/>
            <rect width="190" height="7" transform="matrix(-1 0 0 1 190 0)" fill="#828282"/>
            <rect width="190" height="7" transform="matrix(-1 0 0 1 410 0)" fill="#828282"/>
        </svg>
        <div class="module" style="bottom: 371px" @click="emitModule(1)">
            <Module :id="1" :reverse="false" style=""></Module>
        </div>
        <div class="module" style="bottom: 325px; right: 0px" @click="emitModule(2)">
            <Module :id="2" :reverse="true" style="" ></Module>
        </div>
        <div class="module" style="bottom: 247px;" @click="emitModule(3)">
            <Module :id="3" :reverse="false" style=""></Module>
        </div>
        <div class="module" style="bottom: 191px; right: 0" @click="emitModule(4)">
            <Module :id="4" :reverse="true" style=""></Module>
        </div>
        <div class="module" style="bottom: 123px;" @click="emitModule(5)">
            <Module :id="5" :reverse="false" style="" ></Module>
        </div>
        <div class="module" style="bottom: 77px; right: 0" @click="emitModule(6)">
            <Module :id="6" :reverse="true" style=""></Module>
        </div>
    </div>

</template>

<script>
    import axios from "axios"
    import Module from "./Module";
    export default {
        name: "CatTree",
        components: {
            Module
        },
        methods: {
            emitModule (moduleNumber) {
                console.log(moduleNumber)
                this.$emit('moduleClicked', moduleNumber)
            },
            populateModule(moduleNumber){
                axios.get("http://127.0.0.1:3000/dashboard/cattree/"+moduleNumber.toString())
                    .then(result => {
                        
                        result.data.moduleNumber = moduleNumber.toString()
                        this.$store.commit("FarmUpdate", result.data)
                    })
                    .catch(error => {
                        console.log(error)
                    })
                }
                
            
        },
        created(){
            for(var i = 1;i<7;i++){
                this.populateModule(i)
            }
            
        }
    }
</script>

<style scoped>
    .module{
        position: absolute !important;
    }
</style>