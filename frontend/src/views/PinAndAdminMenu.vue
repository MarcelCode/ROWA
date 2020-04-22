<template>
    <div id="Admin">
        <AdminTopRow v-bind:headtext="'AdminPanel'" v-bind:prevPage="'Home'"></AdminTopRow>
       
        <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'AdminSettings'}">
           <h3>Admin Settings</h3> 
            </v-btn>
         </v-row>
      
        <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'TypeChanger'}">
           <h3>Change Module Plant</h3> 
            </v-btn>
         </v-row>

         <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'BatchFarming'}">
           <h3>Batch Farming</h3> 
            </v-btn>
         </v-row>

         <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'AdminSettings'}">
           <h3>Reality Check</h3> 
            </v-btn>
         </v-row>

         
     
    </div>
</template>


<script>
    
    import AdminTopRow from "@/components/admin/AdminTopRow.vue"
   
    export default {
        name: "AdminMenu",
        components: {
            AdminTopRow,
            
           
        },
         data(){
             return{

             };
         },
         methods: {
              


         },
        
        
        


         
    };
</script>

<style scoped>
    h1,
    h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }

    span {
        color: var(--v-primary-base);
        font-style: normal;
        font-weight: normal;
        font-size: 14px;
    }

    .no-hover:hover {
        background-color: transparent;
        text-decoration: none;
    }

    .info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}
</style>