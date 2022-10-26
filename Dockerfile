FROM webpronl/reveal-md:5.3.4

COPY resources /slides

EXPOSE 8080

CMD [ "/slides", "--theme", "theme/cc.css", "--port", "8080"]