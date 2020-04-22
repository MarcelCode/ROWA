<template>
    <v-btn-toggle rounded class="shadow" borderless>
        <v-btn :style= "[batch_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]"
        @click="change_batch_state">
            Harvest
        </v-btn>
        <v-btn :style= "[!batch_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]"
               @click="change_batch_state">
            Plant
        </v-btn>
    </v-btn-toggle>
</template>

<script>
    import {mapState, mapActions} from "vuex"

    export default {
        name: "BatchButtonPill",
        computed: {
            ...mapState(["batch_active"])
        },
        methods: {
            ...mapActions(["change_batch_state"])
        }
    }
</script>

<style scoped>
    .shadow{box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.25);}
</style>