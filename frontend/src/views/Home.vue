<template>
    <div id="Home">
        <v-container>
            <HomeTopRow/>
            <div v-if="farm_active">
                <v-row style="margin: 30px 0 0 50px">
                    <v-col>
                        <v-chip color="transparent">
                            <v-avatar color="primary" size="16px" style="margin-right: 7px">
                            </v-avatar>
                            <span>Ready to<br>harvest</span>
                        </v-chip>
                        <v-chip color="transparent">
                            <v-avatar color="accent" size="16px" style="margin-right: 7px">
                            </v-avatar>
                            <span>Ready to<br>plant</span>
                        </v-chip>
                    </v-col>
                </v-row>
                <v-row justify="center" style="padding-top: 40px">
                    <CatTree v-on:moduleClicked="zoomToModule"/>
                </v-row>
            </div>
            <div v-else>
                <StatGraphic v-on:infoOn="onInfoOn"></StatGraphic>
            </div>
        </v-container>
         <div>
        <FarmInfo/>
         </div>
    <div>
        <StatInfo v-bind:InfoType="this.info_type_stat"/>
    </div>
         <div>
              <PlantInfo v-bind:InfoType="this.info_type"/> 
            </div>
    </div>
</template>

<script>
    import axios from "axios"
    import HomeTopRow from "@/components/main/HomeTopRow"
    import FarmInfo from "@/components/home/FarmInfo";
    import {mapState} from "vuex";
    import CatTree from "../components/home/Farm/CatTree";
    import StatGraphic from "../components/home/Stats/StatGraphic";
    import PlantInfo from "@/components/home/PlantInfo";
    import StatInfo from "@/components/home/Stats/StatInfo";

    export default {
        name: "Home",
        components: {
            HomeTopRow,
            FarmInfo,
            CatTree,
            StatGraphic,
           PlantInfo,
            StatInfo
        },
        data() {
            return {
                msg: "Farm Overview",
                harvestable_plants: null,
                all_plants: null,
                sensor_data: {
                    datetime: null,
                    temperature: null,
                    light_intensity: null
                },
                sensor_data_updated: null,
                interval: null,
                moduleNumber: null,
                reverse: false,
                info_type: 0,
                info_type_stat: "Undef"
                
            };
        },
        methods: {
             onInfoOn: function (type) {
             this.info_type_stat = type
            this.$store.dispatch('change_ypos_statInfo', ((0+(Math.floor(Math.random() * 10))/1000)));
                
                },
            
            getSensorData: function () {
                this.sensor_data_updated = new Date().toISOString()
                axios.get("http://127.0.0.1:3000/dashboard/sensor-data")
                    .then(result => {
                        this.sensor_data = result.data[0]
                        console.log(this.sensor_data)
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
            getIntervalData: function () {
                this.interval = setInterval(this.getSensorData, 5000);
            },
            zoomToModule: function (moduleNumber) {
                this.moduleNumber = moduleNumber
                this.reverse = moduleNumber % 2 === 0;
                this.info_type = moduleNumber
                this.$store.dispatch('change_ypos_plantInfo', ((-50+(Math.floor(Math.random() * 10)+moduleNumber)/1000)))
            }
        },
        created() {
            // Get sensor data and request them every 10 Seconds
            this.getSensorData()
            // this.getIntervalData()
        },
        beforeDestroy() {
            clearInterval(this.interval)
        },
        computed: {
            ...mapState(["farm_active"])
        }
    };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
</style>
