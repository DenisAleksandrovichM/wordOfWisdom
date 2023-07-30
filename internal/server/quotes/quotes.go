package quotes

import "math/rand"

var quotes = []string{
	"You create your own opportunities. Success doesn’t just come and find you–you have to go out and get it.",
	"Never break your promises. Keep every promise; it makes you credible.",
	"You are never as stuck as you think you are. Success is not final, and failure isn’t fatal.",
	"Happiness is a choice. For every minute you are angry, you lose 60 seconds of your own happiness.",
	"Habits develop into character. Character is the result of our mental attitude and the way we spend our time.",
	"Be happy with who you are. Being happy doesn’t mean everything is perfect but that you have decided to look beyond the imperfections.",
	"Don’t seek happiness–create it. You don’t need life to go your way to be happy.",
	"If you want to be happy, stop complaining. If you want happiness, stop complaining about how your life isn’t what you want and make it into what you do want.",
	"Asking for help is a sign of strength. Don’t let your fear of being judged stop you from asking for help when you need it. Sometimes asking for help is the bravest move you can make. You don’t have to go it alone.",
	"Replace every negative thought with a positive one. A positive mind is stronger than a negative thought.",
	"Accept what is, let go of what was, have faith in what will be. Sometimes you have to let go to let new things come in.",
	"A mind that is stretched by a new experience can never go back to what it was. Experience is what causes a person to make new mistakes instead of old ones.",
	"If you are not willing to learn, no one can help you. If you are determined to learn no one can stop you.",
	"Be confident enough to encourage confidence in others. Show those around you that you have confidence in them.",
	"Allow others to figure things out for themselves. The fixer ends up becoming the enabler. Let people figure it out for themselves; give them a chance to learn.",
	"Confidence is essential for a successful life. Don’t compare yourself to others; compare yourself to the person you were yesterday and give yourself the confidence to be better tomorrow.",
	"Admit your mistakes and don’t repeat them. If you can’t admit your mistakes, you are destined to repeat them.",
	"Be kind to yourself and forgive yourself. You can’t know what you haven’t yet learned, you can’t become proficient without first being a beginner and you can’t be perfect. Welcome to the human race.",
	"Failures are lessons in progress. Failure is always forgivable if you have the courage to learn its lessons and move forward in a new way.",
	"Make amends with those who have wronged you. Apologizing doesn’t always mean that you’re wrong and the other person is right. It just means that you value your relationships more than your ego.",
	"Live your life on your terms. Define what your life is on your own terms and achieve it by your own rules. Build a life you’re proud to live.",
	"When you don’t know, don’t speak as if you do. If you don’t know, simply don’t speak.",
	"Treat others the way you want to be treated. Live by the Golden Rule. Always.",
	"Think before you speak. Never say anything that doesn’t improve on silence.",
	"Cultivate an attitude of gratitude. Never let the things you want make you forget the things you have.",
	"Life isn’t as serious as our minds make it out to be. Life is too short to always be taken so seriously.",
	"Take risks and be bold. At the end, we regret only the chances we didn’t take.",
	"Remember that “no” is a complete sentence. Learn to say no without having to explain yourself.",
	"Don’t feed yourself only on leftovers. When you say yes to others, make sure you are not saying no to yourself.",
	"Build on your strengths. The struggle you are in today is developing the strength you need for tomorrow.",
	"Never doubt your instincts. Trust your hunches; they are usually based on facts filed away in your unconscious mind.",
	"FEAR doesn’t have to stand for Forget Everything and Run. Sometimes it can be Future Expectations Appearing Real.",
	"Your attitude will influence your experience. How you respond is at least as important as what happens to you.",
	"View your life with gentle hindsight. Stop beating yourself up about things from your past. Instead of asking yourself, “What was I thinking?” ask yourself “What was I learning?”",
	"This too shall pass. Just because today is terrible doesn’t mean tomorrow won’t be the best day of your life. You just have to get there.",
}

func GetRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
