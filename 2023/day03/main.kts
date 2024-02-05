import java.io.File

val input = File("input.txt").readText()

val symbols = input.filter { it != '.' && !it.isDigit() && it != '\n' }.toSet()
println(symbols)
val matrix = input.split("\n").map { line -> line.toCharArray() }

// x y
val coordinates =
    listOf(
        -1 to 0,
        1 to 0,
        0 to -1,
        0 to 1,
        -1 to -1,
        -1 to 1,
        1 to -1,
        1 to 1,
    )

data class EngineDigit(val c: Char, val positionX: Int, val positionY: Int)

var sum = 0L

val engineParts = mutableSetOf<List<EngineDigit>>()
for ((y, line) in matrix.withIndex()) {
    for ((x, c) in line.withIndex()) {
        if (c !in symbols) {
            continue
        }

        val matches = mutableSetOf<List<EngineDigit>>()
        for ((dx, dy) in coordinates) {
            var xx = x + dx
            val yy = y + dy
            val coordinateC = matrix.getOrNull(yy)?.getOrNull(xx)
            if (coordinateC?.isDigit() == true) {
                println("found a digit next to $c ($x, $y) -> $coordinateC ($xx, $yy)")
                val startX = xx
                val engineDigits = mutableListOf<EngineDigit>()
                while (true) {
                    val cc = matrix.getOrNull(yy)?.getOrNull(xx)
                    if (cc?.isDigit() == true) {
                        engineDigits.add(0, EngineDigit(cc, xx, yy))
                        xx--
                    } else {
                        break
                    }
                }
                xx = startX + 1
                while (true) {
                    val cc = matrix.getOrNull(yy)?.getOrNull(xx)
                    if (cc?.isDigit() == true) {
                        engineDigits.add(EngineDigit(cc, xx, yy))
                        xx++
                    } else {
                        break
                    }
                }

                matches.add(engineDigits)
            }
        }

        if (c == '*' && matches.size == 2) {
            val (l, r) =
                matches.map { digits ->
                    String(digits.map { it.c }.toCharArray()).toLong()
                }
            sum += l * r
        }

        engineParts.addAll(matches)
    }
}

println(
    engineParts.sumOf { digits ->
        String(digits.map { it.c }.toCharArray()).toInt()
    },
)

println(sum)
