const TelegramBot = require('node-telegram-bot-api');
require('dotenv').config();

const token = process.env.TOKEN;
const bot = new TelegramBot(token, { polling: true });

const conversationSteps = {};

bot.onText(/\/start/, (msg) => {
  const chatId = msg.chat.id;
  const welcomeText = "Привет! Здесь можно подать жалобы если вы считаете что подверглись дискриминации по причине вашего гражданства или узнать какие-то полезные факты. Мы свяжемся с вами для уточнения деталей и вместе подумаем что можно сделать.";
  const options = {
    reply_markup: JSON.stringify({
      inline_keyboard: [
        [{ text: 'Жалоба', callback_data: 'complaint' }]
      ]
    })
  };
  bot.sendMessage(chatId, welcomeText, options);
});

bot.on('callback_query', (callbackQuery) => {
  const chatId = callbackQuery.message.chat.id;
  if (callbackQuery.data === 'complaint') {
    conversationSteps[chatId] = { step: 'askCountry' };
    bot.sendMessage(chatId, "В какой стране случилась проблема?");
  }
});

bot.on('message', (msg) => {  
  const chatId = msg.chat.id;
  if (!conversationSteps[chatId] || msg.text.startsWith('/')) {
    return; 
  }

  if (conversationSteps[chatId].step === 'askCountry') {
    conversationSteps[chatId].country = msg.text;
    conversationSteps[chatId].step = 'askCompany';
    bot.sendMessage(chatId, "Пожалуйста, укажите с какой организацей у Вас возникла проблема.");
    return;
  }

  if (conversationSteps[chatId].step === 'askCompany') {
    conversationSteps[chatId].company = msg.text;
    conversationSteps[chatId].step = 'askType';
    bot.sendMessage(chatId, "Укажите тип организации(банк, авиакомпания и тд.)");
    return;
  }

  if (conversationSteps[chatId].step === 'askType') {
    conversationSteps[chatId].orgType = msg.text;
    conversationSteps[chatId].step = 'askQuestion';
    bot.sendMessage(chatId, "Опишите Вашу проблему подробнее.");
    return;
  }

  if (conversationSteps[chatId].step === 'askQuestion') {
    conversationSteps[chatId].question = msg.text;
    conversationSteps[chatId].step = 'askName';
    bot.sendMessage(chatId, "Как к Вам обращаться?");
    return;
  }


  if (conversationSteps[chatId].step === 'askName') {
    conversationSteps[chatId].name = msg.text;
    conversationSteps[chatId].step = 'askInfo';
    bot.sendMessage(chatId, "Оставьте Ваши контакты для связи(телеграм или электронная почта).");
    return;
  }


  if (conversationSteps[chatId].step === 'askInfo') {
    conversationSteps[chatId].info = msg.text;
    const res = conversationSteps[chatId];

    fetch(`https://geraback.fly.dev:443/issues`, {method: "POST", body: JSON.stringify({
      issuer: {
        name: res.name,
        contact_info: res.info,
      },
      company: {
        country:  res.country,
        name: res.company,
        contact_info: "",
        org_type: res.orgType,
      },
      message: {
        description: res.question
      }
    })}).then((res) => {
      bot.sendMessage(chatId, "Спасибо за Вашу заявку.");
    }).catch((er) => {
      bot.sendMessage(chatId, "ОШИБКА");
    })
    delete conversationSteps[chatId]; 
    return;
  }
});

bot.on('polling_error', (error) => {
  console.log(error);
});
