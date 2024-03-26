const dayjs = require('dayjs')
var relativeTime = require('dayjs/plugin/relativeTime')
dayjs.extend(relativeTime)


export const getDate = (date: string) => {
    return dayjs(date).format("MMM DD, YYYY HH:mm");
}

export const getDateRelative = (date: string) => {
    return dayjs(date).fromNow();
}

export const isOver = (date: string) => {
    return dayjs(date).isBefore(Date.now());
}
