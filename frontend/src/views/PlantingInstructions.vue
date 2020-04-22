<template>
    <div class="container">
        <div class="row">
            <h1>How to plant</h1>
        </div>
        <div class="row">
            <p>Step one etc</p>
        </div>
        <div class="row">
           <router-link to="/">
        <button type="button" class="btn btn-primary" v-on:click="addPlant(pos)">Done</button>
      </router-link>
        </div>
    </div>
</template>


<script>  
import axios from "axios"
export default {
    name: 'PlantingInstructions',
    data (){
        return {
            pos:0
        }
    },
    methods: {
        
        addPlant: function(pos){
            axios.post(
          "http://127.0.0.1:3000/plant/finishedPlanting",
          { PLantedModule: pos },
          "content-type: application/json"
        )
        .then(result => {
          this.position = result.data;
        })
        .catch(error => {
          /*eslint-disable*/
          alert(error);
          /*eslint-enable*/
        });
        },
    },
    created(){
        this.pos = this.$route.params.pos
    }
}
</script>

