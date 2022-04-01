<template>
    <div>
        <v-row v-if="row">
            <v-col v-for="(v, i) in data">
                <v-autocomplete
                    :key="v.Label"
                    :label="v.Label"
                    :items="levelItems(i)"
                    item-text="Name"
                    item-value="ID"
                    v-model="selectedIDs[i]"
                    @change="selectItem($event, i)"
                    :clearable="chips ? false : true"
                    :error-messages="v.ErrorMessages"
                    :chips="chips"
                    :disabled="disabled"
                    :hide-details="hideDetails"
                >
                </v-autocomplete>
            </v-col>
        </v-row>
        <v-autocomplete
            v-else
            v-for="(v, i) in data"
            :key="v.Label"
            :label="v.Label"
            :items="levelItems(i)"
            item-text="Name"
            item-value="ID"
            v-model="selectedIDs[i]"
            @change="selectItem($event, i)"
            :clearable="chips ? false : true"
            :error-messages="v.ErrorMessages"
            :chips="chips"
            :disabled="disabled"
            :hide-details="hideDetails"
        >
        </v-autocomplete>
    </div>
</template>

<script>
export default {
    name: "vx-linkageselect",
    props: {
        value: {
            type: Array,
            default: () => []
        },
        // [{Label, SelectedID, Items: [{ID, Name, ChildrenIDs}], ErrorMessages}]
        data: {
            type: Array,
            default: () => []
        },
        disabled: {
            type: Boolean,
            default: false
        },
        selectOutOfOrder: {
            type: Boolean,
            default: false
        },
        chips: {
            type: Boolean,
            default: false
        },
        row: {
            type: Boolean,
            default: false
        },
        hideDetails: {
            type: Boolean,
            default: false
        },
    },
    data() { 
        return {
            selectedIDs: []
        }
    },
    computed: {
        levelItems: function() {
            return function(level) {
                if (level === 0) {
                    return this.data[0].Items
                }
                if (this.selectedIDs[level-1]) {
                    var idM = {}
                    for (var item of this.data[level-1].Items) {
                        if (item.ID === this.selectedIDs[level-1]) {
                            for (var id of item.ChildrenIDs) {
                                idM[id] = true
                            }
                            break
                        }
                    }
                    var items = []
                    for (var item of this.data[level].Items) {
                        if (idM[item.ID]) {
                            items.push(item)
                        }
                    }
                    return items
                }

                if (this.selectOutOfOrder) {
                    return this.data[level].Items
                }
                return []
            }
        }
    },
    methods: {
        setValue() {
            this.$emit("input", this.selectedIDs)
        },
        validateAndResetSelectedIDs() {
            this.selectedIDs.forEach((v, i) => {
                if (!v) {
                    this.selectedIDs[i] = ""
                    return
                }

                var exists = false
                for (var item of this.data[i].Items) {
                    if (item.ID === v) {
                        exists = true
                        break
                    }
                }
                if (!exists) {
                    this.selectedIDs[i] = ""
                    return
                }

                if (i === 0) {
                    return
                }
                var pID = this.selectedIDs[i-1]
                if (!pID) {
                    if (!this.selectOutOfOrder) {
                        this.selectedIDs[i] = ""
                    }
                    return
                } else {
                    for (var item of this.data[i-1].Items) {
                        if (item.ID === pID) {
                            for (var id of item.ChildrenIDs) {
                                if (id === v) {
                                    return
                                }
                            }
                        }
                    }
                }

                this.selectedIDs[i] = ""
                return
            })
        },
        selectItem(v, level) {
            if (v) {
                if (this.selectedIDs[level+1]) {
                    for (var item of this.data[level].Items) {
                        if (item.ID === v) {
                            var found = false
                            for (var id of item.ChildrenIDs) {
                                if (id === this.selectedIDs[level+1]) {
                                    found = true
                                    break
                                }
                            }
                            if (!found) {
                                for (var i = level+1; i < this.selectedIDs.length; i++) {
                                    this.selectedIDs[i] = ""
                                }
                            }
                        }
                    }
                }
            } else {
                this.selectedIDs[level] = ""
                if (!this.selectOutOfOrder) {
                    for (var i = level+1; i < this.selectedIDs.length; i++) {
                        this.selectedIDs[i] = ""
                    }
                }
            }
            this.setValue()
        }
    },
    mounted() {
        this.data.forEach(e => {
            e.Items.forEach(item => {
                if (!item.Name) {
                    item.Name = item.ID
                }
            })
        })
        if (this.value.length > 0) {
            this.selectedIDs = [...this.value]
        } else {
            this.selectedIDs = this.data.map(e => {
                return e.SelectedID || ""
            })
        }
        this.validateAndResetSelectedIDs()
        this.$nextTick(() => {
            this.setValue();
        });
    }
}
</script>
