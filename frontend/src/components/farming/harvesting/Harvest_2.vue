<template>
<v-container>
<v-row justify="center">
    <h1  style="color:#789659">You will harvest from Module {{moduleNumber}}</h1>
</v-row>
    <v-row justify="center" style="margin-top: 20px">
    <CatTree></CatTree>
     </v-row>

    
    <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
       
   </v-row>
   

    
    </v-container>
</template>

<script>
    import CatTree from "../../home/Farm/CatTree"

    export default {
        name: "harvest_2",
        components: {
            CatTree,
        },
        props: {
            selectedPlant: String,
            moduleNumber: Number
        },
        data() {
            return {}
        },
        methods: {
            getImgUrl(pic) {
                return require('@/assets/harvesting/plants/' + pic + ".png")
            },
        },
        created() {

        }


    }
</script>

<style scoped>

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 20px 100px 0 100px;

    }

    h3 {

        font-family: Montserrat;
        font-style: normal;
        font-weight: 600;
        font-size: 18px;
        line-height: 22px;
        color: var(--v-primary-base)
    }

    #button {
        font-weight: bold;
        margin: 40px;
        font-family: Montserrat;
        font-size: 24px
    }

    h1 {


        font-family: Montserrat;
        font-style: normal;
        font-weight: 600;
        text-align: center;

        color: var(--v-primary-base)
    }
</style>