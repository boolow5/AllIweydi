new Vue({
  el: "#news",
  data: {
    "news_items": [
      {
        "id": 1,
        "title": "Madaafiic hoobiyeyaal oo ku soo dhacay nawaaxiga Xarunta Madaxtooyada",
        "link": "http://www.jowhar.com/2017/03/30/madaafiic-hoobiyeyaal-oo-ku-soo-dhacay-nawaaxiga-xarunta-madaxtooyada/",
        "website_name": "Jowhar",
        "website_url": "/home/mahdi/jowhar.html",
        "created_at": "2017-04-05T19:10:35.878453489Z",
        "updated_at": "2017-04-05T19:10:35.878453607Z"
      },
      {
        "id": 2,
        "title": "Dowladda Itoobiya oo muddo 4 bilood ku kordhisay xukunkii deg dega ahaa",
        "link": "http://www.jowhar.com/2017/03/30/dowladda-itoobiya-oo-muddo-4-bilood-ku-kordhisay-xukunkii-deg-dega-ahaa/",
        "website_name": "Jowhar",
        "website_url": "/home/mahdi/jowhar.html",
        "created_at": "2017-04-05T19:10:37.055414208Z",
        "updated_at": "2017-04-05T19:10:37.055414391Z"
      },
      {
        "id": 3,
        "title": "Madaxweynaha Soomaaliya oo lagu qalday Madaxweynaha Mauritania (Daawo)",
        "link": "http://www.jowhar.com/2017/03/30/madaxweynaha-soomaaliya-oo-lagu-qalday-madaxweynaha-mauritania-daawo/",
        "website_name": "Jowhar",
        "website_url": "/home/mahdi/jowhar.html",
        "created_at": "2017-04-05T19:10:37.18852466Z",
        "updated_at": "2017-04-05T19:10:37.188524828Z"
      },
    ]
  },
  template: `
  <ul class="list-group">
    <li v-for="item in news_items" class="list-group-item">
      Website-ka <a :href="item.website_link">{{item.website_name}}</a> {{timeSince(item.created_at)}} {{$t("before")}} wuxuu qoray:<br />
      <h3><a target="_blank" :href="item.link">{{item.title}}</a></h3>
    </li>
  </ul>`,
  methods: {
    timeSince: function(time) {
      var dateCreated = new Date(time);
      var now = new Date()
      var milliSecondsPassed = now - dateCreated
      var secondsPassed = milliSecondsPassed / (1000)

      var lang = $.cookie('lang');
      var sec = "second(s)";
      var min = "minute(s)";
      var hour = "hour(s)";

      if (lang == "so-SO") {
        sec = "ilbiriqsi"
        min = "daqiiqo"
        hour = "saacadood"
      }

      console.log(lang);

      if (secondsPassed > (60*60)) {
        return Math.round( ((secondsPassed % (60 * 60)))) + " " + hour
      } else if (secondsPassed > 60) {
        return Math.round( ((secondsPassed % 60))) + " " + min
      } else {
        return Math.round((secondsPassed)) + " " + second
      }

    }
  }
})
