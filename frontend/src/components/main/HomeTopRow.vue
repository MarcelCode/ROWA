<template>
    <v-row>
        <v-col cols="4" style="padding-left: 30px">
            <v-btn icon x-large :to="{name:'AdminMenu'}">
                <v-icon>$vuetify.icons.adminSettings</v-icon>
            </v-btn>
        </v-col>
        <v-col align-content-center cols="4">
            <ButtonPill></ButtonPill>
        </v-col>
        <v-col cols="4" id="logo-text" style="text-align: right; padding-right: 30px">
            ROWA
        </v-col>
    </v-row>

</template>

<script>
    import ButtonPill from "./ButtonPill";

    export default {
        name: "HomeTopRow",
        components:{
            ButtonPill
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