const TelegramBot = require('node-telegram-bot-api');
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
    conversationSteps[chatId].step = 'askQuestion';
    bot.sendMessage(chatId, "Какой у вас вопрос?");
    return;
  }

  if (conversationSteps[chatId].step === 'askQuestion') {
    conversationSteps[chatId].question = msg.text;
    console.log(`User ${chatId} from ${conversationSteps[chatId].country} asked: ${conversationSteps[chatId].question}`);
    bot.sendMessage(chatId, "Thank you for your input. We will get back to you soon.");
    delete conversationSteps[chatId]; 
    return;
  }
});

bot.on('polling_error', (error) => {
  console.log(error);
});
