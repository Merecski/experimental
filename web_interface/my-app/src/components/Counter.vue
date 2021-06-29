<template>
    <div id="counter">
        <h1>{{ msg }}</h1>
        Counter: {{ counter }}
    </div> 
</template>

<script>
export default {
    name: "Counter",
    props: {
        msg: String,
        counter: Number,
    },
};
</script>

const Counter = {
    data() {
    return {
        counter: 0
    }
    },
    mounted() {
    setInterval(() => {
        this.counter++
    }, 1000)
    }
}